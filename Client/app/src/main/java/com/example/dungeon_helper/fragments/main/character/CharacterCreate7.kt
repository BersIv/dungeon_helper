package com.example.dungeon_helper.fragments.main.character

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.lifecycle.Observer
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate7Binding
import com.example.dungeon_helper.shared.JsonHelper
import com.example.dungeon_helper.shared.SharedViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import java.io.IOException

class CharacterCreate7 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate7()
    }

    private lateinit var viewModel: CharacterCreate7ViewModel
    private var _binding:FragmentCharacterCreate7Binding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate7ViewModel = ViewModelProvider(this)[CharacterCreate7ViewModel::class.java]
        _binding = FragmentCharacterCreate7Binding.inflate(inflater,container,false)
        val root: View = binding.root
        val textView:TextView = binding.textCharacterCreate7
        characterCreate7ViewModel.text.observe(viewLifecycleOwner)
        {
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()

        val sharedViewModel = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        val goals = binding.goals.editText
        val treasure = binding.treasure.editText
        val notes = binding.notes.editText
        val backBtn = binding.backBtn
        val saveBtn = binding.saveCharBtn


        backBtn.setOnClickListener{
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate7_to_characterCreate6)
        }

        saveBtn.setOnClickListener {

            var newCharacter = sharedViewModel.newCharacter.value

            println(newCharacter)

            newCharacter?.goals = goals?.text.toString()
            newCharacter?.treasures = treasure?.text.toString()
            newCharacter?.notes = notes?.text.toString()

            sharedViewModel.newCharacter.postValue(newCharacter)

            println(sharedViewModel.newCharacter.value)

            sharedViewModel.newCharacter.observe(viewLifecycleOwner, Observer {
                // updating data in displayMsg
                println(it)
                newCharacter = it
            })
//
//
//            sharedViewModel.newCharacter.observe(viewLifecycleOwner, Observer {
//                // updating data in displayMsg
//                println(it)
//                newCharacter = it
//            })
//
//            println(newCharacter)
            sharedViewModel.newCharacter.observe(viewLifecycleOwner, Observer {

                GlobalScope.launch(Dispatchers.Main) {

                    //val serializeCharacter = JsonHelper.serializeToJson(newCharacter)

                    val serializeCharacter = newCharacter?.let { it1 ->
                        JsonHelper.serializeToJson(
                            it1
                        )
                    }

                    println("Cериализованный объект:\n$serializeCharacter")

                    val mediaType = "application/json; charset=utf-8".toMediaType()
                    val body: RequestBody = serializeCharacter.toString().toRequestBody(mediaType)

                    println(body)

                    var token: String = ""
                    sharedViewModel.token.observe(viewLifecycleOwner, Observer {
                        // updating data in displayMsg
                        token = it
                    })


                    //println(getCookieValue())

                    val client = OkHttpClient.Builder().addInterceptor { chain ->
                        val original = chain.request()
                        val authorized = original.newBuilder()
                            .header("Authorization", token)
                            .build()
                        chain.proceed(authorized)
                    }
                        .build()

                    val request = Request.Builder()
                        .url("http://194.247.187.44:5000/character/createCharacter")
                        .post(body)
                        .build()

                    try {
                        val response = withContext(Dispatchers.IO) {
                            client.newCall(request).execute()
                        }

                        if (!response.isSuccessful) {
                            throw IOException("Запрос к серверу не был успешен:" +
                                    " ${response.code} ${response.message}")
                        }

                        println("${response.code} ${response.message}")

                        val activity = requireActivity() as MainActivity
                        val navView = activity.getNavView()
                        val menu = navView.menu
                        val menuItem1 = menu.findItem(R.id.navigation_info)
                        val menuItem2 = menu.findItem(R.id.navigation_character)
                        val menuItem3 = menu.findItem(R.id.navigation_lobby)
                        val menuItem4 = menu.findItem(R.id.navigation_account)
                        menuItem1.isVisible = true
                        menuItem2.isVisible = true
                        menuItem3.isVisible = true
                        menuItem4.isVisible = true


                        (activity as MainActivity).navController.navigate(R.id.action_characterCreate7_to_navigation_character)

                    } catch (e: IOException) {
                        println("Ошибка подключения: $e");
                    }
                }

            })


        }
    }
}