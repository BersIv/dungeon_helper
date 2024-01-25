package com.example.dungeon_helper.fragments.main.character

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.AdapterView
import android.widget.ArrayAdapter
import android.widget.TextView
import android.widget.Toast
import androidx.lifecycle.Observer
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate3Binding
import com.example.dungeon_helper.shared.JsonHelper
import com.example.dungeon_helper.shared.Race
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
import org.json.JSONObject
import java.io.IOException


class CharacterCreate3 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate3()
    }

    private lateinit var viewModel: CharacterCreate3ViewModel
    private var _binding: FragmentCharacterCreate3Binding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate3ViewModel = ViewModelProvider(this)[CharacterCreate3ViewModel::class.java]
        _binding = FragmentCharacterCreate3Binding.inflate(inflater,container,false)
        val root: View = binding.root
        val textView:TextView = binding.textCharacterCreate3
        characterCreate3ViewModel.text.observe(viewLifecycleOwner){
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

        val optionsSpinner = binding.spinnerOptions
        val backBtn = binding.backBtn
        val forwBtn = binding.forwBtn
        var selected = ""

        var listRaces = sharedViewModel.races.value
        sharedViewModel.races.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            listRaces = it
        })


        val arrayOptions: ArrayList<String> = ArrayList()

        if (listRaces != null) {
            for (i in listRaces!!) {
                arrayOptions.add(i.raceName)
            }
        }


        val optionsArray = arrayOf("Гном", "Дварф", "Драконорожденный", "Полуорк", "Полурослик","Полуэльф","Тифлинг","Человек",
            "Эльф")

        val adapter =
            ArrayAdapter(requireContext(), android.R.layout.simple_spinner_item, arrayOptions)

        adapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item)


        optionsSpinner.adapter = adapter


        optionsSpinner.onItemSelectedListener = object : AdapterView.OnItemSelectedListener {
            override fun onItemSelected(
                parentView: AdapterView<*>?,
                selectedItemView: View?,
                position: Int,
                id: Long
            ) {

                val selectedItem = parentView?.getItemAtPosition(position).toString()
                Toast.makeText(requireContext(), "Выбрано: $selectedItem", Toast.LENGTH_SHORT)
                    .show()

                selected = selectedItem
            }

            override fun onNothingSelected(parentView: AdapterView<*>?) {
                forwBtn.isEnabled = false

            }

        }


        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate3_to_characterCreate2)
        }

        forwBtn.setOnClickListener{
            GlobalScope.launch(Dispatchers.Main) {

                var raceId = Race(0, "")

                if (listRaces != null) {
                    for (i in listRaces!!) {
                        if (selected == i.raceName) {
                            raceId = i
                        }
                    }
                }

                val jsonBody = JSONObject().apply {
                    put("idRace", raceId.id)
                }

                println(jsonBody)

                val mediaType = "application/json; charset=utf-8".toMediaType()
                val body: RequestBody = jsonBody.toString().toRequestBody(mediaType)

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
                    .url("http://194.247.187.44:5000/subrace/getSubraces")
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

                    val deserializedSubraceList = JsonHelper.deserializeListFromJsonSubraceList(response.body!!.string())
                    println("Десериализованный объект:\n$deserializedSubraceList")

                    sharedViewModel.subraces.value = deserializedSubraceList

                    var race = Race(0, "")

                    if (listRaces != null) {
                        for (i in listRaces!!) {
                            if (selected == i.raceName) {
                                race = i
                            }
                        }
                    }

                    val newCharacter = sharedViewModel.newCharacter.value

                    newCharacter?.race = race

                    (activity as MainActivity).navController.navigate(R.id.action_characterCreate3_to_characterCreate4)

                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }
            }


        }



    }



}