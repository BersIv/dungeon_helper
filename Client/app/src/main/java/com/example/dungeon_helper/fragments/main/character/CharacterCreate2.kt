package com.example.dungeon_helper.fragments.main.character

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.TextView
import android.widget.AdapterView
import android.widget.Toast
import androidx.lifecycle.Observer
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate2Binding
import com.example.dungeon_helper.shared.CharacterClass
import com.example.dungeon_helper.shared.JsonHelper
import com.example.dungeon_helper.shared.SharedViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.OkHttpClient
import okhttp3.Request
import java.io.IOException

class CharacterCreate2 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate2()
    }

    private lateinit var viewModel: CharacterCreate2ViewModel
    private var _binding: FragmentCharacterCreate2Binding? = null
    private val binding get() = _binding!!


    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate2ViewModel = ViewModelProvider(this)[CharacterCreate2ViewModel::class.java]
        _binding =  FragmentCharacterCreate2Binding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textCharacterCreate2
        characterCreate2ViewModel.text.observe(viewLifecycleOwner){
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

        var listCharacterClass = sharedViewModel.characterClasses.value

        sharedViewModel.characterClasses.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            listCharacterClass = it
        })



        val arrayOptions: ArrayList<String> = ArrayList()

        if (listCharacterClass != null) {
            for (i in listCharacterClass!!) {
                arrayOptions.add(i.className)
            }
        }

        val optionsArray = arrayOf("Бард", "Варвар", "Бард", "Воин", "Волшебник","Друид","Жрец",
            "Палладин","Чародей", "Следопыт","Плут","Колдун","Монах")

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
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate2_to_characterCreate1)
        }

        forwBtn.setOnClickListener {
            GlobalScope.launch(Dispatchers.Main) {

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
                    .url("http://194.247.187.44:5000/race/getRaces")
                    .get()
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

                    val deserializedRacesList = JsonHelper.deserializeListFromJsonRaceList(response.body!!.string())
                    println("Десериализованный объект:\n$deserializedRacesList")

                    sharedViewModel.races.value = deserializedRacesList

                    var characterClass = CharacterClass(0, "")

                    if (listCharacterClass != null) {
                        for (i in listCharacterClass!!) {
                            if (selected == i.className) {
                                characterClass = i
                            }
                        }
                    }

                    val newCharacter = sharedViewModel.newCharacter.value

                    newCharacter?.charClass = characterClass

                    (activity as MainActivity).navController.navigate(R.id.action_characterCreate2_to_characterCreate3)


                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }
            }



        }


    }

}