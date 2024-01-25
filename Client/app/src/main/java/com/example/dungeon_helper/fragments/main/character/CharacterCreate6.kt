package com.example.dungeon_helper.fragments.main.character

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate6Binding
import com.example.dungeon_helper.shared.SharedViewModel

class CharacterCreate6 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate6()
    }

    private lateinit var viewModel: CharacterCreate6ViewModel
    private var _binding: FragmentCharacterCreate6Binding? = null
    private val binding get() = _binding!!


    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate6ViewModel = ViewModelProvider(this)[CharacterCreate6ViewModel::class.java]
        _binding = FragmentCharacterCreate6Binding.inflate(inflater,container,false)
        val root: View = binding.root
        val textView:TextView = binding.textCharacterCreate6
        characterCreate6ViewModel.text.observe(viewLifecycleOwner)
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

        val ideals = binding.ideals.editText
        val weakneses = binding.weakness.editText
        val traits = binding.characterTraits.editText
        val allies = binding.allies.editText
        val organizations = binding.organizations.editText
        val enemies = binding.enemies.editText
        val story = binding.story.editText
        val backBtn = binding.backBtn
        val forwBtn = binding.forwBtn

        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate6_to_characterCreate5)
        }

        forwBtn.setOnClickListener {

            val newCharacter = sharedViewModel.newCharacter.value

            newCharacter?.ideals = ideals?.text.toString()
            newCharacter?.weaknesses = weakneses?.text.toString()
            newCharacter?.traits = traits?.text.toString()
            newCharacter?.allies = allies?.text.toString()
            newCharacter?.organizations = organizations?.text.toString()
            newCharacter?.enemies = enemies?.text.toString()
            newCharacter?.story = story?.text.toString()


            (activity as MainActivity).navController.navigate(R.id.action_characterCreate6_to_characterCreate7)
        }
    }
}